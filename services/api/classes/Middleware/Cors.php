<?php
declare(strict_types=1);

namespace Middleware;

use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Server\RequestHandlerInterface as RequestHandler;
use Slim\Psr7\Response;

class Cors
{
	public function __invoke(ServerRequestInterface $request, RequestHandler $handler): ResponseInterface
	{
		if ($request->getMethod() !== 'OPTIONS') {
			$response = $handler->handle($request);
		} else {
			$response = new Response();
		}

		$response = $response->withHeader('Access-Control-Allow-Origin', $this->getOrigin($request));
		$response = $response->withHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS');
		$response = $response->withHeader('Access-Control-Allow-Headers', $request->getHeaderLine('Access-Control-Request-Headers'));
		$response = $response->withHeader('Access-Control-Allow-Credentials', 'true');

		return $response;
	}

	private function getOrigin(ServerRequestInterface $request): string
	{
		if ($origin = $request->getHeaderLine('HTTP_ORIGIN')) {
			return $origin;
		}

		if ($origin = $request->getHeaderLine('HTTP_REFERER')) {
			return $origin;
		}

		return '*';
	}
}
