<script lang="ts">
	import { Code2, LayoutGrid, Lock, ShieldCheck, Sparkles, SquareStack } from '@lucide/svelte';

	let {
		features = []
	}: {
		features: {
			title: string;
			description: string;
			icon: string;
		}[];
	} = $props();

	const iconMap = {
		shield: ShieldCheck,
		layout: LayoutGrid,
		code: Code2,
		lock: Lock,
		stack: SquareStack,
		spark: Sparkles
	} as const;
</script>

<div class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
	{#each features as feature}
		{@const Icon = iconMap[feature.icon as keyof typeof iconMap]}
		<article class="group rounded-[26px] border border-[#e7e3d8] bg-white/85 p-5 shadow-[0_12px_30px_rgba(21,31,23,0.05)] transition hover:-translate-y-0.5 hover:shadow-[0_16px_34px_rgba(21,31,23,0.08)]">
			<div class="flex items-start justify-between gap-4">
				<div class="grid h-11 w-11 place-items-center rounded-2xl bg-[#dcefe3] text-[#145336]">
					<Icon size={18} strokeWidth={2.1} />
				</div>
				<span class="rounded-full bg-[#f5f3ed] px-3 py-1 text-[11px] font-semibold uppercase tracking-[0.24em] text-[#7a847d]">
					{feature.icon}
				</span>
			</div>
			<h3 class="mt-5 text-base font-semibold tracking-tight text-[#162117]">{feature.title}</h3>
			<p class="mt-2 text-sm leading-6 text-[#445046]">{feature.description}</p>
		</article>
	{/each}
</div>
