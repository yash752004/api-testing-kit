<section class="grid gap-4 xl:grid-cols-[0.95fr_1.05fr]">
	<div class="rounded-[28px] border border-[#e7e3d8] bg-white/90 p-6 shadow-[0_12px_30px_rgba(21,31,23,0.05)]">
		<p class="text-xs font-semibold uppercase tracking-[0.3em] text-[#1f7a4d]">Architecture</p>
		<h2 class="mt-2 text-2xl font-semibold tracking-tight text-[#162117] sm:text-3xl">
			A stack that is straightforward to ship, explain, and extend
		</h2>
		<p class="mt-3 text-sm leading-6 text-[#445046] sm:text-base">
			The repo docs already point to SvelteKit on the front end, Go for the HTTP API, PostgreSQL for persistence, and Goose for migrations. The page keeps that structure visible instead of hiding it behind hype.
		</p>

		<div class="mt-5 grid gap-3 sm:grid-cols-2">
			<div class="rounded-2xl border border-[#dfe8dd] bg-[#f5f3ed] px-4 py-3">
				<p class="text-xs font-semibold uppercase tracking-[0.22em] text-[#7a847d]">Frontend</p>
				<p class="mt-1 text-sm font-medium text-[#162117]">SvelteKit, Tailwind, reusable sections</p>
			</div>
			<div class="rounded-2xl border border-[#dfe8dd] bg-[#f5f3ed] px-4 py-3">
				<p class="text-xs font-semibold uppercase tracking-[0.22em] text-[#7a847d]">Backend</p>
				<p class="mt-1 text-sm font-medium text-[#162117]">Go HTTP API with explicit routing</p>
			</div>
			<div class="rounded-2xl border border-[#dfe8dd] bg-[#f5f3ed] px-4 py-3">
				<p class="text-xs font-semibold uppercase tracking-[0.22em] text-[#7a847d]">Data</p>
				<p class="mt-1 text-sm font-medium text-[#162117]">PostgreSQL, sessions, history, and templates</p>
			</div>
			<div class="rounded-2xl border border-[#dfe8dd] bg-[#f5f3ed] px-4 py-3">
				<p class="text-xs font-semibold uppercase tracking-[0.22em] text-[#7a847d]">Local dev</p>
				<p class="mt-1 text-sm font-medium text-[#162117]">Docker-friendly setup and test scripts</p>
			</div>
		</div>
	</div>

	<div class="rounded-[28px] border border-[#e7e3d8] bg-[linear-gradient(180deg,rgba(31,122,77,0.12),rgba(255,255,255,0.94))] p-6 shadow-[0_12px_30px_rgba(21,31,23,0.05)]">
		<p class="text-xs font-semibold uppercase tracking-[0.3em] text-[#1f7a4d]">Safety model</p>
		<h2 class="mt-2 text-2xl font-semibold tracking-tight text-[#162117] sm:text-3xl">
			The app is designed to feel open without becoming an open relay
		</h2>
		<p class="mt-3 text-sm leading-6 text-[#445046] sm:text-base">
			The product direction in the docs is explicit: guests can only use allowlisted examples, and authenticated requests still pass destination validation, redirect checks, and conservative limits.
		</p>

		<div class="mt-5 space-y-3">
			<div class="rounded-2xl border border-white/75 bg-white/80 px-4 py-3 text-sm text-[#445046]">Block localhost, private ranges, and metadata targets.</div>
			<div class="rounded-2xl border border-white/75 bg-white/80 px-4 py-3 text-sm text-[#445046]">Validate redirects and keep response previews bounded.</div>
			<div class="rounded-2xl border border-white/75 bg-white/80 px-4 py-3 text-sm text-[#445046]">Log suspicious attempts for later admin review.</div>
		</div>
	</div>
</section>
