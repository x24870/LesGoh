GOERLI_INTERCHAIN_ACCOUNT_ROUTER=0x55486284a85d7b51a7bBfd343702414D65276fa6
# GOERLI_IGP=0x8f9C3888bFC8a5B25AED115A82eCbb788b196d2a
GOERLI_DEFUALT_IGP=0xF90cB82a76492614D07B82a7658917f3aC811Ac1
# OP GOERLI
# DESTINATION=420
# MUMBAI
DESTINATION=80001
# ARB GOERLI
# DESTINATION=421613
RECIPIENT=0x36FdA966CfffF8a9Cdc814f546db0e6378bFef35
# calldata: "fooBar(uint256,string)" 666 "DAT ASS"
CALLDATA=0xf07c1f47000000000000000000000000000000000000000000000000000000000000029a000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000074441542041535300000000000000000000000000000000000000000000000000
RPC_URL=https://ethereum-goerli.publicnode.com		
MY_ADDRESS=0x3C127Be7c30e4f0c863047ca566684d66a264893
PRIVATE_KEY=407cd3d457a111e79784cf2bbbad73dd512eaac2e5b0492be113fe11222ef98c


send:
	cast send ${GOERLI_INTERCHAIN_ACCOUNT_ROUTER} \
'callRemote(uint32,address,uint256,bytes)' \
${DESTINATION} ${RECIPIENT} 0 ${CALLDATA} \
--rpc-url ${RPC_URL} --private-key ${PRIVATE_KEY}

GAS_AMOUNT=550000

qoute:
	cast call ${GOERLI_DEFUALT_IGP} "quoteGasPayment(uint32,uint256)" ${DESTINATION} ${GAS_AMOUNT} --rpc-url ${RPC_URL}

GAS_PAYMENT_QUOTE=0x000000000000000000000000000000000000000000000000001d805cfd528800
MESSAGE_ID=0x7cb64f872f565cb564a29f272d02408fc7cf5fbdbc388195d561674bb84aee94


pay:
	cast send ${GOERLI_DEFUALT_IGP} "payForGas(bytes32,uint32,uint256,address)" \
${MESSAGE_ID} ${DESTINATION} ${GAS_AMOUNT} ${MY_ADDRESS} --rpc-url ${RPC_URL} \
--private-key ${PRIVATE_KEY} --value ${GAS_PAYMENT_QUOTE}