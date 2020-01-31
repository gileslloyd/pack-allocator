<?php
declare(strict_types=1);

namespace Response;

class Resultset implements ToArrayInterface
{

	use ToArrayTrait;

	/**
	 * @var int
	 */
	protected $count;

	/**
	 * @var int
	 */
	protected $total;

	/**
	 * @var int
	 */
	protected $offset;

	/**
	 * @var int
	 */
	protected $limit;

	public function __construct(int $count, int $total, int $offset = 0, int $limit = 0)
	{
		$this->count = $count;
		$this->total = $total;
		$this->offset = $offset;
		$this->limit = $limit;
	}

	public function getCount(): int
	{
		return $this->count;
	}

	public function getTotal(): int
	{
		return $this->total;
	}

	public function getOffset(): int
	{
		return $this->offset;
	}

	public function getLimit(): int
	{
		return $this->limit;
	}

}
