FAUCET_URL=http://0.0.0.0:8081
REST_URL=http://0.0.0.0:8080
TESTNET_URL=https://fullnode.testnet.aptoslabs.com
OWNER=0x0
TESTNET=0x56260c419e8b176e0ca7f6d439b69180c2de2cb284d8dee24476f247af204492

init_profiles:
	aptos init --profile owner --rest-url ${REST_URL} --faucet-url ${FAUCET_URL}
	aptos init --profile user --rest-url ${REST_URL} --faucet-url ${FAUCET_URL}

init_testnet_prof:
	aptos init --profile testnet --rest-url "https://fullnode.testnet.aptoslabs.com/v1"

local_testnet:
	aptos node run-local-testnet --with-faucet

local_testnet_restart:
	aptos node run-local-testnet --with-faucet --force-restart

test:
	aptos move test --skip-fetch-latest-git-deps --named-addresses owner=owner 

fund:
	aptos account fund-with-faucet \
	--profile owner --account owner --amount 999999999
	aptos account fund-with-faucet \
	--profile user --account user --amount 999999999

compile:
	aptos move compile --bytecode-version 6 --named-addresses owner=${profile}

compile_testnet:
	aptos move compile --named-addresses owner=testnet

publish:
	aptos move publish --named-addresses owner=${profile} \
	--bytecode-version=6 --included-artifacts=none \
	--sender-account ${profile} --profile ${profile}

publish_testnet:
	aptos move publish --named-addresses owner=testnet \
	--sender-account testnet --profile testnet

mint_shovel:
	aptos move run --function-id ${TESTNET}::urn_to_earn::mint_shovel \
	--sender-account=${profile} --profile=${profile}
	
wl_mint_shovel:
	aptos move run --function-id ${OWNER}::urn_to_earn::bayc_wl_mint_shovel \
	--sender-account=user --profile=user

add_burned:
	aptos move run --function-id ${OWNER}::urn::add_burned \
	--sender-account=owner --profile=owner \
	--args address:${OWNER} bool:false

mint_shovel_testnet:
	aptos move run-script --assume-yes \
	--compiled-script-path build/urn/bytecode_scripts/mint_shovel.mv \
	--sender-account=testnet --profile=testnet

mint_urn_testnet:
	aptos move run-script --assume-yes \
	--compiled-script-path build/urn/bytecode_scripts/mint_urn.mv \
	--sender-account=testnet --profile=testnet

dig:
	aptos move run --function-id ${OWNER}::urn_to_earn::dig \
	--sender-account=owner --profile=owner

dig_testnet:
	aptos move run-script --assume-yes \
	--compiled-script-path build/urn/bytecode_scripts/dig.mv \
	--sender-account=testnet --profile=testnet

query_owner_res:
		aptos account list --query resources --account owner --profile owner

query_user_res:
		aptos account list --query resources --account user --profile user

query_testnet_res:
		aptos account list --query resources --account testnet --profile testnet

