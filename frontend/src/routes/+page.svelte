<script lang="ts">
	interface PdfsDataProp {
		name: string
		url: string
	}

	let pdfsData: PdfsDataProp[] | null = $state(null)
	let searchQuery = $state("")

	let apiUrl = import.meta.env.DEV ? "http://localhost:8000" : "https://apipdfs.vercel.app"

	$effect(() => {
		fetch(apiUrl)
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
	<div class="space-y-2 mb-10 flex flex-col items-center">
		<h1 class="text-5xl font-bold">PDFs</h1>
		<p 
		class="text-lg"
		>
		If you want to contribute,
		<a 
		href="https://github.com/srcrapi/pdfs-site"
		target="_blank"
		class="text-slate-400 hover:text-secondary"
		>
		click here
		</a>
		and fork
		</p>
	</div>

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
					href={pdf.url}
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
