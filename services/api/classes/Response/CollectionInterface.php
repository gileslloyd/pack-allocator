<?php
declare(strict_types=1);

namespace Response;

interface CollectionInterface extends ToArrayInterface
{
	public function getCount(): int;
	public function getTotal(): int;
	public function getOffset(): int;
	public function getLimit(): int;
	public function getItems(): array;
}
