<?php

declare(strict_types=1);

namespace Controller;

use Response\SuccessSingleResponse;
use Slim\Psr7\Request;
use Slim\Psr7\Response;
use Response\ErrorDetails;
use Response\ErrorResponse;

class PackAllocationController extends \Controller
{
	public function get(Request $request, Response $response, array $args): Response
	{
		try {
			var_dump($request); die;

			$api_response = new SuccessSingleResponse($packAllocation, 'Successful Pack Allocation');
		} catch (\Exception $e) {
			$response = $response->withStatus(400);
			$api_response = new ErrorResponse('Failed to calculate pack allocation', 400);
			$exception_string = get_class($e) . "[{$e->getFile()}:{$e->getLine()}]";
			$api_response->addError(new ErrorDetails('400', $e->getMessage(), $exception_string));
		}

		return $this->renderer->render($request, $response, $api_response->toArray());
	}

	private function validateInput(Request $request): int
	{
		if (!isset($request->getQueryParams()['items'])) {
			throw new \Exception('You must provide the number of items required');
		}

		$requiredItems = (int) $request->getQueryParams()['items'];
		if ($requiredItems < 1) {
			throw new \Exception('The number of items required must be a positive number');
		}

		return $requiredItems;
	}
}
