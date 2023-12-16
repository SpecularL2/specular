// SPDX-License-Identifier: MIT
pragma solidity ^0.8.11;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {StandardBridge} from "./StandardBridge.sol";
import {L1Portal} from "./L1Portal.sol";
import {IMintableERC20} from "./mintable/IMintableERC20.sol";
import {Predeploys} from "../libraries/Predeploys.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";

contract L1StandardBridge is StandardBridge {
    using SafeERC20 for IERC20;

    L1Portal public L1_PORTAL;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @inheritdoc StandardBridge
    modifier onlyOtherBridge() override {
        address origSender = L1_PORTAL.l2Sender();
        require(
            msg.sender == address(PORTAL_ADDRESS) && origSender == address(OTHER_BRIDGE),
            "StandardBridge: function can only be called from the other bridge"
        );

        _;
    }

    /// @notice Initializer;
    function initialize(address payable _l1Portal) public initializer {
        L1_PORTAL = L1Portal(_l1Portal);
        __StandardBridge_init(_l1Portal, payable(Predeploys.L2_STANDARD_BRIDGE));
    }

    function _authorizeUpgrade(address) internal override onlyOwner whenPaused {}

    /// @inheritdoc StandardBridge
    receive() external payable override whenNotPaused {
        _initiateBridgeETH(msg.sender, msg.sender, msg.value, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }

    /// @inheritdoc StandardBridge
    function _initiateBridgeETH(
        address _from,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes memory _extraData
    ) internal override {
        emit ETHBridgeInitiated(_from, _to, _amount, _extraData);

        L1_PORTAL.initiateDeposit{value: _amount}(
            address(OTHER_BRIDGE),
            _minGasLimit,
            abi.encodeCall(this.finalizeBridgeETH, (_from, _to, _amount, _extraData))
        );
    }

    /// @inheritdoc StandardBridge
    function _initiateBridgeERC20(
        address _localToken,
        address _remoteToken,
        address _from,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes memory _extraData
    ) internal override {
        if (_isNonNativeTokenPair(_localToken, _remoteToken)) {
            IMintableERC20(_localToken).burn(_from, _amount);
        } else {
            IERC20(_localToken).safeTransferFrom(_from, address(this), _amount);
            deposits[_localToken][_remoteToken] = deposits[_localToken][_remoteToken] + _amount;
        }

        emit ERC20BridgeInitiated(_localToken, _remoteToken, _from, _to, _amount, _extraData);

        L1_PORTAL.initiateDeposit(
            address(OTHER_BRIDGE),
            _minGasLimit,
            abi.encodeWithSelector(
                this.finalizeBridgeERC20.selector,
                // Because this call will be executed on the remote chain,
                // we reverse the order of
                // the remote and local token addresses relative to their order in the
                // finalizeBridgeERC20 function.
                _remoteToken,
                _localToken,
                _from,
                _to,
                _amount,
                _extraData
            )
        );
    }
}
