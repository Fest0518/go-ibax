/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
	(next_id('1_system_parameters'),'default_ecosystem_menu', '', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'default_ecosystem_contract', '', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'gap_between_blocks', '2', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'rollback_blocks', '60', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'honor_nodes', '', 'ContractAccess("@1UpdateSysParam","@1NodeRemoveByKey")'),
	(next_id('1_system_parameters'),'number_of_nodes', '101', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_ecosystem', '10000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_table', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_column', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_contract', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_menu', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_page', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_block', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_view', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_application', '10000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_token', '50000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_asset', '10000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_block_size', '67108864', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_tx_size', '33554432', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_tx_block', '1000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_columns', '50', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_indexes', '5', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_tx_block_per_user', '1000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_fuel_tx', '20000000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_fuel_block', '200000000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'taxes_size', '3', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'taxes_wallet', '', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'fuel_rate', '[["1","1000000"]]', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_address_to_id', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_id_to_address', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_hash', '50', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_pub_to_id', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_ecosys_param', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_sys_param_string', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_sys_param_int', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_sys_fuel', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_validate_condition', '30', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_eval_condition', '20', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_has_prefix', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_contains', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_replace', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_join', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_update_lang', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_size', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_substr', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_contracts_list', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_is_object', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_compile_contract', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_flush_contract', '50', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_eval', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_len', '5', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_bind_wallet', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_unbind_wallet', '10', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_create_ecosystem', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_table_conditions', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_create_table', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_perm_table', '100', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_column_condition', '50', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_create_column', '50', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_perm_column', '50', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_json_to_map', '50', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'max_block_generation_time', '2000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'block_reward','10','ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'incorrect_blocks_per_day','10','ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'node_ban_time','86400000','ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'node_ban_time_local','1800000','ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_tx_size_wallet', '15', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_create_rate', '1000000', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'test','false','false'),
	(next_id('1_system_parameters'),'price_tx_data', '0', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_contract_by_name', '20', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'price_exec_contract_by_id', '20', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'private_blockchain', '1', 'false'),
	(next_id('1_system_parameters'),'external_blockchain', '', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'pay_free_contract', '@1CallDelayedContract', 'ContractAccess("@1UpdateSysParam")'),
	(next_id('1_system_parameters'),'pool_block_rate', '5', 'ContractAccess("@1UpdateSysParam")'),
    (next_id('1_system_parameters'),'local_node_ban_time', '60', 'ContractAccess("@1UpdateSysParam")');
`
