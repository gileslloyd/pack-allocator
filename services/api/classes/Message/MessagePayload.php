<?php

declare(strict_types=1);

namespace Message;

use Response\ToArrayInterface;

class MessagePayload implements ToArrayInterface
{
	/**
	 * @var array
	 */
	private $payload = [];

	public function __construct(array $payload)
	{
		$this->payload = $payload;
	}

	public function getPayload(): array
	{
		return $this->payload;
	}

	public function get(string $key, string $default = null)
	{
		$value = $this->payload;

		foreach (explode('.', $key) as $part) {
			if (!array_key_exists($part, $value)) {
				return $default;
			}

			$value = $value[$part];
		}

		return $value;
	}

	public function toArray(): array
	{
		return $this->getPayload();
	}
}
