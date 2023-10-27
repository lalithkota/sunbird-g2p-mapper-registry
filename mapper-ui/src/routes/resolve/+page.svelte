<script>
	import {createMapperHeader, createTransactionObject} from "../../util.js";
	import {config} from "../../config.js";

	let idValue = "";
	let faValue = "";

	function onResolve() {
		const {transaction_id, reference_id} = createTransactionObject();
		fetch(config.BASE_URL + "/mapper-service/v0.1.0/mapper/resolve", {
			method: 'post',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				"signature": "Signature:  namespace=\"g2p\", kidId=\"{sender_id}|{unique_key_id}|{algorithm}\", algorithm=\"ed25519\", created=\"1606970629\", expires=\"1607030629\", headers=\"(created) (expires) digest\", signature=\"Base64(signing content)",
				"header": createMapperHeader("resolve"),
				"message": {
					"transaction_id": transaction_id,
					"resolve_request": [
						{
							"reference_id": reference_id,
							"timestamp": new Date().toISOString(),
							"fa": faValue,
							"id": idValue,
							"name": "Mr. John Smith",
							"scope": idValue && faValue ? "yes_no" : "details",
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
		<input bind:value={idValue}/>
	</div>
	<div>
		<span>FA:</span>
		<input bind:value={faValue}/>
	</div>

	<button on:click={onResolve}>Resolve</button>
</div>
<style>
	div {
		margin-top: 1rem;
	}
</style>
