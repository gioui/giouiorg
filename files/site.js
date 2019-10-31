$(function() {
	$("[data-run=browser]").each(function() {
		$(this).attr("content-editable", "true");
		var cont = $(this).parent();
		var run = $('<button>Run</button>');
		var output = $('<div></div>');
		$(cont).append(run);
		$(cont).append(output);
		playground({
			codeEl: $(this),
			outputEl: output,
			runEl: run
		});
	})
})

