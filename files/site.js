$(function() {
	'use strict';

	$("[data-run=browser]").each(function() {
		var code = this;
		$(code).attr("contenteditable", "true");
		$(code).attr("spellcheck", "false");
		var cont = $(code).parent();
		var run = $('<button>Run</button>');
		$(cont).append(run);
		var output = $('<div></div>');
		output.hide();
		$(cont).append(output);
		var transport = new HTTPTransport();
		function onRun() {
			output.text("");
			output.show();
			transport.Run($(code).text(), PlaygroundOutput(output.get(0)), null);
		}
		run.click(onRun);
	})
})

