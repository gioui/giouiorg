$(function() {
	'use strict';

	$("[data-run=playground]").each(function() {
		var code = this;
		$(code).attr("contenteditable", "true");
		$(code).attr("spellcheck", "false");
		var cont = $(code).parent();
		$(cont).addClass("play");
		var run = $('<button class="run">Run</button>');
		$(cont).append(run);
		var output = $('<pre class="output"></pre>');
		output.hide();
		$(output).insertAfter(cont);
		var transport = new HTTPTransport();
		function onRun() {
			$(output).text("");
			$(output).show();
			var program = $(code).text();
			program = program.replace('\xA0', ' '); // replace non-breaking spaces
			transport.Run(program, PlaygroundOutput(output.get(0)), null);
		}
		run.click(onRun);
	})
	$("[data-run=wasm]").each(function() {
		var code = this;
		var cont = $(code).parent();
		$(cont).addClass("play");
		var run = $('<button class="run">Run</button>');
		$(cont).append(run);
		var args = $(code).data("args")
		var win = $('<div class="window"><iframe src="/wasm/'+args+'"></iframe></div>');
		function onRun() {
			$(win).insertAfter(cont);
		}
		run.click(onRun);
	})
})

