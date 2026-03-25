<script lang="ts">
	import { ArrowRight, FileCode2, PanelTop, Users } from '@lucide/svelte';

	let {
		useCases = []
	}: {
		useCases: {
			title: string;
			text: string;
			steps: string[];
		}[];
	} = $props();

	const icons = [PanelTop, FileCode2, Users];
</script>

<div class="grid gap-4 xl:grid-cols-3">
	{#each useCases as useCase, index}
		{@const Icon = icons[index]}
		<article class="rounded-[28px] border border-[#e7e3d8] bg-[linear-gradient(180deg,rgba(255,255,255,0.92),rgba(245,243,237,0.96))] p-5 shadow-[0_12px_30px_rgba(21,31,23,0.05)]">
			<div class="flex items-start justify-between gap-3">
				<div class="grid h-11 w-11 place-items-center rounded-2xl bg-[#1f7a4d] text-white shadow-[0_10px_24px_rgba(31,122,77,0.22)]">
					<Icon size={18} strokeWidth={2.1} />
				</div>
				<span class="rounded-full border border-[#dce7d8] bg-white px-3 py-1 text-[11px] font-semibold uppercase tracking-[0.22em] text-[#1f7a4d]">
					0{index + 1}
				</span>
			</div>

			<h3 class="mt-5 text-base font-semibold tracking-tight text-[#162117]">{useCase.title}</h3>
			<p class="mt-2 text-sm leading-6 text-[#445046]">{useCase.text}</p>

			<ul class="mt-4 space-y-2">
				{#each useCase.steps as step}
					<li class="flex items-start gap-2 text-sm leading-6 text-[#445046]">
						<ArrowRight size={15} class="mt-1 shrink-0 text-[#1f7a4d]" />
						<span>{step}</span>
					</li>
				{/each}
			</ul>
		</article>
	{/each}
</div>
