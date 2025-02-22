// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IBridgeAdapter} from "../interfaces/IBridge.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

interface IArbSys {
  function withdrawEth(address destination) external payable returns (uint256);
}

interface IL2GatewayRouter {
  function outboundTransfer(
    address l1Token,
    address to,
    uint256 amount,
    bytes calldata data
  ) external payable returns (bytes memory);
}

/// @notice Arbitrum L2 Bridge adapter
/// @dev Auto unwraps and re-wraps wrapped eth in the bridge.
contract ArbitrumL2BridgeAdapter is IBridgeAdapter {
  using SafeERC20 for IERC20;

  IL2GatewayRouter internal immutable i_l2GatewayRouter;
  //  address internal immutable i_l1ERC20Gateway;
  IArbSys internal constant ARB_SYS = IArbSys(address(0x64));

  constructor(IL2GatewayRouter l2GatewayRouter) {
    if (address(l2GatewayRouter) == address(0)) {
      revert BridgeAddressCannotBeZero();
    }
    i_l2GatewayRouter = l2GatewayRouter;
  }

  /// @inheritdoc IBridgeAdapter
  function sendERC20(
    address localToken,
    address remoteToken,
    address recipient,
    uint256 amount
  ) external payable override returns (bytes memory) {
    if (msg.value != 0) {
      revert MsgShouldNotContainValue(msg.value);
    }

    IERC20(localToken).safeTransferFrom(msg.sender, address(this), amount);

    // No approval needed, the bridge will burn the tokens from this contract.
    // TODO: return data bombs?
    return i_l2GatewayRouter.outboundTransfer(remoteToken, recipient, amount, bytes(""));
  }

  /// @notice No-op since L1 -> L2 transfers do not need finalization.
  function finalizeWithdrawERC20(
    address /* remoteSender */,
    address /* localReceiver */,
    bytes calldata /* bridgeSpecificPayload */
  ) external {}

  /// @notice There are no fees to bridge back to L1
  function getBridgeFeeInNative() external pure returns (uint256) {
    return 0;
  }

  function depositNativeToL1(address recipient) external payable {
    ARB_SYS.withdrawEth{value: msg.value}(recipient);
  }
}
