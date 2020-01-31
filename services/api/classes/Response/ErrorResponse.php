<?php
declare(strict_types=1);

namespace Response;

class ErrorResponse extends BaseResponse
{

	/**
	 * @var array
	 */
	protected $errors = [];

	public function __construct(string $message, int $code = 200, string $details = '')
	{
		$meta = new Meta('error', $code, $message, $details);

		parent::__construct($meta);
	}

	public function addError(ErrorDetails $error)
	{
		array_push($this->errors, $error);
	}

}
