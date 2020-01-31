<?php
declare(strict_types=1);

namespace Response;

class CollectionMeta extends Meta
{

	/**
	 * @var Resultset
	 */
	protected $resultset;

	public function setResultset(Resultset $resultset)
	{
		$this->resultset = $resultset;
	}

}
