<?php
require_once('plugins/login-servers.php');

return new AdminerLoginServers(
	$servers = array(
		'Postgres' => array('server' => 'postgres', 'driver' => 'pgsql'),
	)
);