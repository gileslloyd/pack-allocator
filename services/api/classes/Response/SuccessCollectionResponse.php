<?php
declare(strict_types=1);

namespace Response;

class SuccessCollectionResponse extends BaseResponse
{

	/**
	 * @var CollectionInterface
	 */
	protected $data;

	public function __construct(CollectionInterface $data, string $message, string $details = '', int $code = 200)
	{
		$this->data = $data;

		$meta = new CollectionMeta('success', $code, $message, $details);

		$meta->setResultset(
			new Resultset(
				$this->data->getCount(),
				$this->data->getTotal(),
				$this->data->getOffset(),
				$this->data->getLimit()
			)
		);

		parent::__construct($meta);
	}

}
