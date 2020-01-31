<?php
declare(strict_types=1);

namespace Middleware;

use Psr\Http\Message\ServerRequestInterface as Request;
use Psr\Http\Server\RequestHandlerInterface as RequestHandler;
use Slim\Psr7\Response;

class Auth
{
	private $userId;

	public function __invoke(Request $request, RequestHandler $handler): Response
	{
		if(!$this->isAuthorised($request->getHeader("User-Token"))){
			return new Response(401);
		}

		$request = $request->withAttribute('userId', $this->userId);
		$response = $handler->handle($request);

		return $response;
	}

	private function isAuthorised(array $headers): bool
	{
		if(!reset($headers)){
			return false;
		}

		if (!$this->tokenIsValid(reset($headers))) {
			return false;
		}

		return true;
	}

	private function tokenIsValid(string $token)
	{
		$messageBus = \DI\Container::instance()->get('message_bus');

		$response = $messageBus->sync(
			'fmf-auth',
			[
				'role' => 'token',
				'cmd' => 'validate',
				'payload' => [
					'token' => $token,
				],
			]
		);

		if ($response->get('isValid')) {
			$this->userId = $response->get('id');
			return true;
		}

		return false;
	}
}