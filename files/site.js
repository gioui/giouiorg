$(function() {
	'use strict';

	// preserve sidebar scroll position between page changes.
	var sidebar = document.querySelector(".layout-sidebar");
	var top = sessionStorage.getItem("sidebar-scroll");
	if (top !== null) {
		sidebar.scrollTop = parseInt(top, 10);
	}
	window.addEventListener("beforeunload", function() {
		sessionStorage.setItem("sidebar-scroll", sidebar.scrollTop);
	});

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
		var pkg = $(this).data("pkg");
		var size = $(this).data("size");
		var args = $(this).data("args");
		var width = "";
		var height = "";
		if (size !== undefined) {
			var spl = size.split("x");
			width = spl[0];
			if (spl.length == 2) {
				height = spl[1];
			}
		}
		var cont = $(this);
		var replace = true;
		// The markdown formatter puts the args on the inner <code>
		// element for {{/path/code.go}} tags.
		if ($(cont).prop("tagName") == "CODE") {
			// For code snippets the wasm program should display after
			// the snippet.
			replace = false;
			cont = $(cont).parent();
		}
		$(cont).addClass("play");
		var run = $('<button class="run">Run</button>');
		$(cont).append(run);
		var src = "/files/wasm/"+pkg+"/index.html";
		if (args) {
			src = src + "?argv=" + encodeURIComponent(args);
		}
		var win = $('<div class="window"><iframe width="'+width+'" height="'+height+'" src="'+src+'"></iframe></div>');
		function onRun() {
			if (replace) {
				$(cont).empty();
				$(cont).append(win);
			} else {
				var p = $('<p></p>');
				$(p).append(win);
				$(p).insertAfter(cont);
			}
		}
		run.click(onRun);
	})

	// Current workaround for the issue https://todo.sr.ht/~eliasnaur/gio/415
	//  where keydown and keyup listeners don't work
	window.addEventListener('keydown', (e)=> {
		const frames = document.querySelectorAll('iframe')
		frames.forEach((frame)=> {
			const input = frame.contentDocument.querySelector('input')
			if (input) {
				input.dispatchEvent(new KeyboardEvent('keydown', {key:e.key}))
			}
		})
	});
	window.addEventListener('keyup', (e)=> {
		const frames = document.querySelectorAll('iframe')
		frames.forEach((frame)=> {
			const input = frame.contentDocument.querySelector('input')
			if (input) {
				input.dispatchEvent(new KeyboardEvent('keyup', {key:e.key}))
			}
		})
	});
})

