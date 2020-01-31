<?php
declare(strict_types=1);

namespace Response;

trait ToArrayTrait
{

	/**
	 * Recursively convert an object into an array
	 *
	 * @throws \UnexpectedValueException If a child object does not implement the necessary ToArrayInterface
	 *
	 * @return array
	 */
	public function toArray(): array
	{
		return array_map([$this, 'toArrayValue'], get_object_vars($this));
	}

	public function toArrayValue($value)
	{
		switch ($type = gettype($value)) {
			case 'boolean':
			case 'integer':
			case 'float':
			case 'double':
			case 'string':
			case 'NULL':
				return $value;

			case 'array':
				return array_map([$this, 'toArrayValue'], $value);

			case 'object':
				if ($value instanceof \DateTimeInterface) {
					// ISO 8601 date format (2004-02-12T15:19:21+00:00)
					return $value->format('c');
				}

				if ($value instanceof ToArrayInterface) {
					return $value->toArray();
				}
				throw new \UnexpectedValueException('Objects must implement ' . ToArrayInterface::class);

			default:
				throw new \UnexpectedValueException("Value of type '{$type}' can not be converted");
		}
	}
}
