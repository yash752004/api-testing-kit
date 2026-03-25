<script lang="ts">
	const response = String.raw`{
  "name": "sveltejs/kit",
  "stars": 18800,
  "safety": "validated"
}`;
</script>

<section class="space-y-5">

	<div class="max-w-3xl">
		<p class="text-xs font-semibold uppercase tracking-[0.3em] text-[#1f7a4d]">Workspace preview</p>
		<h2 class="mt-2 text-2xl font-semibold tracking-tight text-[#162117] sm:text-3xl">
			The request builder and response viewer stay visible together
		</h2>
		<p class="mt-3 text-sm leading-6 text-[#445046] sm:text-base">
			Guests can inspect curated examples, while signed-in users get the full workspace with saved requests, collections, and stricter outbound control.
		</p>
	</div>

	<div class="relative">
		<div class="absolute inset-0 -z-10 rounded-[32px] bg-[radial-gradient(circle_at_top_right,rgba(31,122,77,0.2),transparent_52%)]"></div>
		<div class="rounded-[32px] border border-[#e7e3d8] bg-[linear-gradient(180deg,rgba(255,255,255,0.9),rgba(245,243,237,0.92))] p-4 shadow-[0_18px_40px_rgba(21,31,23,0.08)] sm:p-5">
			<div class="rounded-[28px] border border-[#e7e3d8] bg-white p-4">
				<div class="flex items-center justify-between border-b border-[#eef0ea] pb-4">
					<div>
						<p class="text-xs font-semibold uppercase tracking-[0.24em] text-[#1f7a4d]">Live preview</p>
						<p class="mt-1 text-sm text-[#7a847d]">Curated request + response canvas</p>
					</div>
					<div class="rounded-full bg-[#dcefe3] px-3 py-1 text-xs font-semibold text-[#145336]">Guest mode</div>
				</div>

				<div class="mt-4 grid gap-3 lg:grid-cols-[0.92fr_1.08fr]">
					<div class="rounded-[24px] border border-[#e7e3d8] bg-[#f5f3ed] p-4">
						<div class="flex items-center justify-between">
							<p class="text-sm font-semibold text-[#162117]">Request</p>
							<span class="rounded-full bg-[#dcefe3] px-2.5 py-1 text-xs font-semibold text-[#145336]">GET</span>
						</div>
						<div class="mt-3 rounded-2xl border border-[#dfe8dd] bg-white px-3 py-2 text-xs font-mono text-[#445046]">
							https://api.github.com/repos/sveltejs/kit
						</div>
						<div class="mt-3 space-y-2">
							<div class="flex items-center justify-between rounded-2xl bg-white px-3 py-2 text-xs text-[#445046]">
								<span>Headers</span><span>4 items</span>
							</div>
							<div class="flex items-center justify-between rounded-2xl bg-white px-3 py-2 text-xs text-[#445046]">
								<span>Body</span><span>Locked for this example</span>
							</div>
							<div class="flex items-center justify-between rounded-2xl bg-white px-3 py-2 text-xs text-[#445046]">
								<span>Auth</span><span>None</span>
							</div>
						</div>
					</div>

					<div class="rounded-[24px] border border-[#e7e3d8] bg-[#162117] p-4 text-white shadow-[0_12px_30px_rgba(21,31,23,0.16)]">
						<div class="flex items-center justify-between border-b border-white/10 pb-4">
							<div>
								<p class="text-xs font-semibold uppercase tracking-[0.24em] text-[#9bc4a8]">Response</p>
								<p class="mt-1 text-sm text-white/65">Pretty JSON with metadata</p>
							</div>
							<div class="rounded-full bg-[#1f7a4d] px-3 py-1 text-xs font-semibold">200 OK</div>
						</div>

						<div class="mt-4 grid gap-3 sm:grid-cols-3">
							<div class="rounded-2xl bg-white/6 px-3 py-3">
								<p class="text-[11px] uppercase tracking-[0.22em] text-white/50">Time</p>
								<p class="mt-1 text-sm font-semibold">183 ms</p>
							</div>
							<div class="rounded-2xl bg-white/6 px-3 py-3">
								<p class="text-[11px] uppercase tracking-[0.22em] text-white/50">Size</p>
								<p class="mt-1 text-sm font-semibold">18.2 KB</p>
							</div>
							<div class="rounded-2xl bg-white/6 px-3 py-3">
								<p class="text-[11px] uppercase tracking-[0.22em] text-white/50">Type</p>
								<p class="mt-1 text-sm font-semibold">application/json</p>
							</div>
						</div>

						<div class="mt-4 overflow-hidden rounded-[22px] border border-white/10 bg-[#0f1712]">
							<div class="flex gap-3 border-b border-white/10 px-4 py-3 text-xs font-semibold text-white/55">
								<span class="text-white">Pretty</span>
								<span>Raw</span>
								<span>Headers</span>
							</div>
							<pre class="overflow-x-auto px-4 py-4 text-xs leading-6 text-[#d8ebdf]"><code>{response}</code></pre>
						</div>
					</div>
				</div>

				<div class="mt-4 rounded-[24px] border border-[#e7e3d8] bg-[#f5f3ed] px-4 py-3">
					<div class="flex flex-wrap items-center justify-between gap-3">
						<p class="text-sm font-semibold text-[#162117]">Snippet drawer</p>
						<div class="flex flex-wrap gap-2 text-xs">
							<span class="rounded-full bg-white px-3 py-1 text-[#445046]">curl</span>
							<span class="rounded-full bg-white px-3 py-1 text-[#445046]">fetch</span>
							<span class="rounded-full bg-white px-3 py-1 text-[#445046]">python requests</span>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>
