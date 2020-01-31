<?php

declare(strict_types=1);

use RKA\ContentTypeRenderer\Renderer;
use Message\MessageBus;

abstract class Controller
{
	/**
	 * @var MessageBus
	 */
	protected $message_bus;

	/**
	 * @var \RKA\ContentTypeRenderer\Renderer
	 */
	protected $renderer;

	public function __construct()
	{
		$this->message_bus = new MessageBus(
			getenv('RABBIT_HOST'),
			(int) getenv('RABBIT_PORT'),
			getenv('RABBIT_USER'),
			getenv('RABBIT_PASSWORD')
		);

		$this->renderer = new Renderer();
	}
}
