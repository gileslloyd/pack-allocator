<?php
declare(strict_types=1);

namespace Response;

class FailResponse extends BaseResponse
{

	/**
	 * @var array
	 */
	protected $errors = [];

	public function __construct(string $message, int $code = 200, string $details = '')
	{
		$meta = new Meta('fail', $code, $message, $details);

		parent::__construct($meta);
	}

	public function addError(FailDetails $error)
	{
		array_push($this->errors, $error);
	}

}
