<script>
	let id = "";
	let fa = "";
	let name = "";
	let mobile_number = "";
	function onUpdate() {
		fetch("https://g2p-financial-wrapper.xiv.in/mapper-service/v0.1.0/mapper/update", {
			method: 'post',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				"signature": "Signature:  namespace=\"g2p\", kidId=\"{sender_id}|{unique_key_id}|{algorithm}\", algorithm=\"ed25519\", created=\"1606970629\", expires=\"1607030629\", headers=\"(created) (expires) digest\", signature=\"Base64(signing content)",
				"header": {
					"version": "0.1.0",
					"message_id": "123456789020211216223812",
					"message_ts": "2022-12-04T18:01:07+00:00",
					"action": "update",
					"sender_id": "registry.example.org",
					"sender_uri": "https://registry.sender.org/g2p/callback/on-disburse",
					"receiver_id": "pymts.example.org",
					"total_count": 21800,
					"is_encrypted": true
				},
				"message": {
					"transaction_id": "12345678901234567000",
					"update_request": [
						{
							"reference_id": "12345678901234567890",
							"timestamp": "2022-12-04T17:20:07-04:00",
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
	<button on:click={onUpdate}>Update</button>
</div>
<style>
	div {
		margin-top: 1rem;
	}
</style>
