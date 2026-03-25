<script lang="ts">
	import HeroSection from '$lib/components/marketing/hero-section.svelte';
	import ProductPreview from '$lib/components/marketing/product-preview.svelte';
	import FeatureGrid from '$lib/components/marketing/feature-grid.svelte';
	import UseCases from '$lib/components/marketing/use-cases.svelte';
	import CredibilitySection from '$lib/components/marketing/credibility-section.svelte';
	import SiteFooter from '$lib/components/marketing/site-footer.svelte';

	const features = [
		{
			title: 'Guest-safe request exploration',
			description:
				'Visitors can try curated examples immediately while the backend keeps arbitrary outbound targets locked down.',
			icon: 'shield'
		},
		{
			title: 'Split request and response workspace',
			description:
				'The core canvas keeps method, URL, headers, body, and response metadata visible in one working surface.',
			icon: 'layout'
		},
		{
			title: 'Response inspection that feels technical',
			description:
				'Pretty, raw, and headers views stay readable with monospace surfaces, status badges, and concise metadata.',
			icon: 'code'
		},
		{
			title: 'Safety-first outbound execution',
			description:
				'Private ranges, metadata IPs, and unsupported protocols stay blocked so the product never behaves like an open proxy.',
			icon: 'lock'
		},
		{
			title: 'Reusable collections and history',
			description:
				'Authenticated users get save flows, re-run paths, and durable request history without changing the app model.',
			icon: 'stack'
		},
		{
			title: 'Built for polished demos',
			description:
				'The landing page and workspace share the same warm, premium visual language so the product feels coherent.',
			icon: 'spark'
		}
	];

	const useCases = [
		{
			title: 'Portfolio reviews',
			text: 'Show a real working product with guest-safe examples, clear constraints, and a visible upgrade path.',
			steps: ['Open the app', 'Run a sample request', 'Inspect the response panel']
		},
		{
			title: 'API debugging',
			text: 'Use the same workspace for method changes, header edits, response validation, and quick retries.',
			steps: ['Tweak the request', 'Send it again', 'Compare responses side by side']
		},
		{
			title: 'Team onboarding',
			text: 'Keep a lightweight, teachable surface for introducing APIs without the overhead of a full platform.',
			steps: ['Start with templates', 'Save useful requests', 'Reuse them in collections']
		}
	];

</script>

<svelte:head>
	<title>API Testing Kit - Safe API workspace for demos and real requests</title>
	<meta
		name="description"
		content="A warm, premium API testing workspace with guest-safe examples, strong outbound controls, and a polished request-response experience."
	/>
</svelte:head>

<div class="relative isolate overflow-hidden bg-[radial-gradient(circle_at_top_left,_rgba(31,122,77,0.16),_transparent_32%),linear-gradient(180deg,_#f2f0ea_0%,_#ece8df_100%)] text-[#162117]">
	<div class="pointer-events-none absolute inset-0 overflow-hidden">
		<div class="absolute -left-32 top-24 h-80 w-80 rounded-full bg-[#1f7a4d]/10 blur-3xl"></div>
		<div class="absolute right-[-4rem] top-1/3 h-96 w-96 rounded-full bg-[#dcefe3] blur-3xl"></div>
		<div class="absolute bottom-[-8rem] left-1/2 h-72 w-72 -translate-x-1/2 rounded-full bg-[#145336]/10 blur-3xl"></div>
	</div>

	<div class="mx-auto min-h-screen max-w-[1440px] px-4 py-4 sm:px-6 lg:px-8">
		<div class="overflow-hidden rounded-[32px] border border-[#e7e3d8] bg-[rgba(247,245,240,0.9)] shadow-[0_24px_60px_rgba(21,31,23,0.08)] backdrop-blur">
			<header class="border-b border-[#e7e3d8] bg-white/75 px-5 py-4 sm:px-6 lg:px-8">
				<div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
					<div class="flex items-center gap-3">
						<div class="grid h-11 w-11 place-items-center rounded-2xl bg-[#1f7a4d] text-sm font-semibold text-white shadow-[0_10px_24px_rgba(31,122,77,0.28)]">
							AT
						</div>
						<div>
							<p class="text-sm font-semibold tracking-tight text-[#162117]">API Testing Kit</p>
							<p class="text-xs text-[#7a847d]">A safer way to demo and ship HTTP workflows</p>
						</div>
					</div>

					<nav class="flex flex-wrap items-center gap-2 text-sm">
						<a href="#features" class="rounded-full px-4 py-2 text-[#445046] transition hover:bg-[#f5f3ed] hover:text-[#162117]">Features</a>
						<a href="#examples" class="rounded-full px-4 py-2 text-[#445046] transition hover:bg-[#f5f3ed] hover:text-[#162117]">Examples</a>
						<a href="#trust" class="rounded-full px-4 py-2 text-[#445046] transition hover:bg-[#f5f3ed] hover:text-[#162117]">Safety</a>
						<a
							href="/app"
							class="rounded-full bg-[#1f7a4d] px-5 py-2.5 font-medium text-white shadow-[0_12px_28px_rgba(31,122,77,0.28)] transition hover:bg-[#19663f]"
						>
							Open app
						</a>
					</nav>
				</div>
			</header>

			<main class="space-y-12 px-5 py-6 sm:px-6 lg:px-8 lg:py-8">
				<HeroSection />
				<ProductPreview />

				<section id="features" class="scroll-mt-24 space-y-6">
					<div class="max-w-3xl">
						<p class="text-xs font-semibold uppercase tracking-[0.3em] text-[#1f7a4d]">Why it works</p>
						<h2 class="mt-2 text-2xl font-semibold tracking-tight text-[#162117] sm:text-3xl">
							A product surface that looks polished before a user signs in
						</h2>
						<p class="mt-3 text-sm leading-6 text-[#445046] sm:text-base">
							The landing page mirrors the app's tone: warm, premium, and technically credible. Guests can explore the product, while the backend keeps execution constraints explicit.
						</p>
					</div>

					<FeatureGrid {features} />
				</section>

				<section id="examples" class="scroll-mt-24 space-y-6">
					<div class="max-w-3xl">
						<p class="text-xs font-semibold uppercase tracking-[0.3em] text-[#1f7a4d]">Examples</p>
						<h2 class="mt-2 text-2xl font-semibold tracking-tight text-[#162117] sm:text-3xl">
							Use cases that show the product, not just the copy
						</h2>
					</div>

					<UseCases {useCases} />
				</section>

				<section id="trust" class="scroll-mt-24">
					<CredibilitySection />
				</section>

				<section class="rounded-[30px] border border-[#dfe8dd] bg-[linear-gradient(135deg,rgba(31,122,77,0.16),rgba(255,255,255,0.96))] p-6 shadow-[0_18px_40px_rgba(21,31,23,0.08)] sm:p-8">
					<div class="flex flex-col gap-5 lg:flex-row lg:items-center lg:justify-between">
						<div class="max-w-2xl">
							<p class="text-xs font-semibold uppercase tracking-[0.3em] text-[#1f7a4d]">Ready to inspect it</p>
							<h2 class="mt-2 text-2xl font-semibold tracking-tight text-[#162117] sm:text-3xl">
								Open the app and start with a safe demo request
							</h2>
							<p class="mt-3 text-sm leading-6 text-[#445046] sm:text-base">
								The workspace is built to feel real immediately: request builder, response viewer, and clear locked states for anything the guest mode should not do.
							</p>
						</div>

						<div class="flex flex-wrap gap-3">
							<a
								href="/app"
								class="rounded-full bg-[#1f7a4d] px-6 py-3 font-medium text-white shadow-[0_12px_28px_rgba(31,122,77,0.28)] transition hover:bg-[#19663f]"
							>
								Open /app
							</a>
							<a
								href="#features"
								class="rounded-full border border-[#d4ded1] bg-white px-6 py-3 font-medium text-[#162117] transition hover:bg-[#f5f3ed]"
							>
								Review the surface
							</a>
						</div>
					</div>
				</section>
			</main>

			<SiteFooter />
		</div>
	</div>
</div>
