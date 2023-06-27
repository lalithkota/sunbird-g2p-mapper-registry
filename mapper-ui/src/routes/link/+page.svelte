<script>
	import {createMapperHeader, createTransactionObject} from "../../util.js";
	import {config} from "../../config.js";

	let id = "";
	let fa = "";
	let name = "";
	let mobile_number = "";
	function onLink() {
		const {transaction_id, reference_id} = createTransactionObject();
		fetch(config.BASE_URL + "/mapper-service/v0.1.0/mapper/link", {
			method: 'post',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				"signature": "Signature:  namespace=\"g2p\", kidId=\"{sender_id}|{unique_key_id}|{algorithm}\", algorithm=\"ed25519\", created=\"1606970629\", expires=\"1607030629\", headers=\"(created) (expires) digest\", signature=\"Base64(signing content)",
				"header": createMapperHeader("link"),
				"message": {
					"transaction_id": transaction_id,
					"link_request": [
						{
							"reference_id": reference_id,
							"timestamp": new Date().toISOString(),
							"id": id,
							"fa": fa,
							"name": name,
							"mobile_number": mobile_number,
							"additional_info": {
								"key": "string",
								"value": "string"
							},
							"locale": "eng"
						}
					]
				}
			})
		})
				.then(response => response.json())
				.then(data => {
					console.log(data);
				}).catch(error => {
			console.log(error);
			return [];
		});
	}
</script>
<div class="text-column">
	<div>
		<span>ID:</span>
		<input bind:value={id}/>
	</div>
	<div>
		<span>FA:</span>
		<input bind:value={fa}/>
	</div>
	<div>
		<span>Name:</span>
		<input bind:value={name}/>
	</div>
	<div>
		<span>Mobile Number:</span>
		<input bind:value={mobile_number}/>
	</div>
	<button on:click={onLink}>Link</button>
</div>
<style>
	div {
		margin-top: 1rem;
	}
</style>
