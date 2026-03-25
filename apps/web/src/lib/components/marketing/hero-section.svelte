<script lang="ts">
	const stats = [
		{ label: 'Guest-safe endpoints', value: 'Curated' },
		{ label: 'Response viewer', value: 'Pretty + raw' },
		{ label: 'Backend controls', value: 'Validated' }
	];
</script>

<section class="space-y-6">
	<div class="space-y-6">
		<div class="inline-flex items-center gap-2 rounded-full border border-[#d9e7d8] bg-white/80 px-4 py-2 text-xs font-semibold uppercase tracking-[0.28em] text-[#1f7a4d]">
			<span class="h-2 w-2 rounded-full bg-[#1f7a4d]"></span>
			Safe demo workspace
		</div>

		<div class="space-y-4">
			<h1 class="max-w-3xl text-4xl font-semibold tracking-tight text-[#162117] sm:text-5xl lg:text-6xl">
				A premium API testing workspace that is safe to put in front of people
			</h1>
			<p class="max-w-2xl text-sm leading-7 text-[#445046] sm:text-base">
				API Testing Kit combines a guest-friendly landing experience with a real request/response surface, explicit execution controls, and a backend that refuses to behave like an open proxy.
			</p>
		</div>

		<div class="flex flex-wrap gap-3">
			<a
				href="/app"
				class="rounded-full bg-[#1f7a4d] px-6 py-3 font-medium text-white shadow-[0_12px_28px_rgba(31,122,77,0.28)] transition hover:bg-[#19663f]"
			>
				Open the app
			</a>
			<a
				href="#examples"
				class="rounded-full border border-[#d4ded1] bg-white px-6 py-3 font-medium text-[#162117] transition hover:bg-[#f5f3ed]"
			>
				See example flows
			</a>
		</div>

		<div class="grid gap-3 sm:grid-cols-3">
			{#each stats as stat}
				<div class="rounded-2xl border border-[#e7e3d8] bg-white/80 p-4 shadow-[0_10px_24px_rgba(21,31,23,0.04)]">
					<p class="text-xs font-medium uppercase tracking-[0.22em] text-[#7a847d]">{stat.label}</p>
					<p class="mt-2 text-base font-semibold text-[#162117]">{stat.value}</p>
				</div>
			{/each}
		</div>
	</div>
</section>
