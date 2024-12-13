<script lang="ts">
	interface PdfsDataProp {
		name: string
		download_url: string
	}

	let pdfsData: PdfsDataProp[] | null = $state(null)

	$effect(() => {
		fetch("http://localhost:8000")
		.then(resp => {
			if (!resp.ok) throw new Error("failed to get response")

			return resp.json()
		})
		.then(pdfData => {
			pdfsData = pdfData
		})
		.catch(err => console.error(err))
	})
</script>

<main class="text-white p-4 space-y-2 text-lg">
{#if pdfsData}
		{#each pdfsData as pdf}
				<a href={pdf.download_url}>{pdf.name.split(".")[0]}</a> <br>
		{/each}
{:else}
	<p>Loading...</p>
{/if}
</main>
