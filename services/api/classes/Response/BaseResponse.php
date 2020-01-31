<?php
declare(strict_types=1);

namespace Response;

abstract class BaseResponse implements ToArrayInterface
{

	use ToArrayTrait;

	/**
	 * @var Meta
	 */
	protected $meta;

	public function __construct(Meta $meta)
	{
		$this->meta = $meta;
	}

}
