<script lang="ts">
	interface PdfsDataProp {
		name: string
		download_url: string
	}

	let pdfsData: PdfsDataProp[] | null = $state(null)
	let searchQuery = $state("")

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

	let filteredPdfs = $derived.by(() => {
		if (!pdfsData) return []

		return pdfsData.filter(pdf =>
			pdf.name.toLowerCase().includes(searchQuery.toLowerCase()))
	})
</script>

<main class="h-svh flex flex-col items-center justify-center space-y-5 text-white">
{#if pdfsData}
	<h1 class="text-5xl font-bold mb-10">PDFs</h1>

	<input 
	type="text" 
	placeholder="Search..." 
	bind:value={searchQuery}
	class="bg-zinc-700 rounded-xl p-1.5"
	>

	<div class="bg-zinc-800 rounded-xl w-[66rem] h-[32rem] p-3 mx-52 overflow-y-auto space-y-2 text-lg">
		{#if filteredPdfs.length > 0}
			{#each filteredPdfs as pdf}
				<p class="p-1 hover:bg-zinc-900 hover:rounded-md">
					<a
					href={pdf.download_url}
					>
					{pdf.name.split(".")[0]}
					</a>
					<br>
				</p>
			{/each}
		{:else}
			<p class="text-center text-xl">No results found</p>
		{/if}
	</div>
{:else}
	<p class="text-center text-2xl font-bold">Loading...</p>
{/if}
</main>
