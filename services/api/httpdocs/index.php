<?php

ini_set('display_errors', 1);
error_reporting(-1);

$autoloader = require '../vendor/autoload.php';

$app = \Slim\Factory\AppFactory::create();

require '../config/routes.php';

$app->add(new \Middleware\Cors());

$app->run();
