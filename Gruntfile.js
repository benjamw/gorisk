/* jshint esversion: 6 */

module.exports = function (grunt) {
	"use strict";

	grunt.initConfig({

		handlebars: {
			templates: {
				files: {
					"static/js/compiledHandlebars/compiledHandlebarsTemplates.js": "static/handlebars/templates/**/*.hbs"
				}
			},
			partials: {
				files: {
					"static/js/compiledHandlebars/compiledHandlebarsPartials.js": "static/handlebars/partials/**/*.hbs"
				}
			}
		},

		concat: {
			javascript: {
				files: {
					"static/js/combined.js": [
						"static/js/vendor/jquery1.12.4.min.js", // jQuery first
						"static/js/vendor/handlebars.runtime-v4.0.5.js", // then Handlebars

						"static/js/compiledHandlebars/compiledHandlebarsPartials.js",
						"static/js/compiledHandlebars/compiledHandlebarsTemplates.js",

						// grab everything else
						"static/js/**/*.js",

						// remove combined, main, and init
						"!static/js/combined.js",
						"!static/js/main.js",
						"!static/js/init.js",

						// add back in main, so it's second to last
						"static/js/main.js",

						// add back in init, so it's last
						"static/js/init.js"
					]
				}
			}
		},

		sass: {
			all: {
				options: {
					sourcemap: "none",
					style: "expanded", // Output style. Can be nested, compact, compressed, expanded.
					unixNewlines: true, // Force Unix newlines in written files.
					update: true, // Only compile changed files
					lineNumbers: false // Emit comments in the generated CSS indicating the corresponding source line.
				},
				files: {
					"static/css/main.css": "static/scss/main.scss"
				}
				// files: [{
				//     expand: true,
				//     cwd: "static/scss/",
				//     src: ["*.scss"],
				//     dest: "static/css/",
				//     ext: ".css",
				// }]
			}
		},

		watch: {
			handlebarsTemplates: {
				files: ["static/handlebars/templates/**/*.hbs"],
				tasks: ["handlebars:templates"]
			},
			handlebarsPartials: {
				files: ["static/handlebars/partials/**/*.hbs"],
				tasks: ["handlebars:partials"]
			},
			javascript: {
				files: ["static/js/**/*.js", "!static/js/combined.js"],
				tasks: ["concat:javascript"]
			},
			sass: {
				files: ["static/scss/**/*.scss"],
				tasks: ["sass"]
			},
			group: {
				files: ["static/css/**/*.css", "static/js/combined.js"],
				tasks: ["group"]
			},
			gruntfile: {
				files: ["Gruntfile.js"],
				tasks: ["default"]
			}
		}

	});

	// Load Task Plugins
	grunt.loadNpmTasks("grunt-contrib-handlebars");
	grunt.loadNpmTasks("grunt-contrib-concat");
	grunt.loadNpmTasks("grunt-contrib-sass");
	grunt.loadNpmTasks("grunt-contrib-watch");

	// Group
	grunt.registerTask("group", "Rebuild css/js groupings", vdotRegroup);

	// Default Tasks
	grunt.registerTask("default", ["handlebars", "concat", "sass", "group"]);


	function vdotRegroup() {
		var minifyGroups = true;
		var groupInName = "static/group_in_file.json";
		var groupOutName = "static/group_out_file.js";

		// Let's import some classes
		var fs = require("fs");
		var md5 = require("blueimp-md5");
		var json = require("comment-json");
		var shell = require("shelljs");

		if (minifyGroups) {
			var jsmin = require("jsmin").jsmin;
			var cssmin = require("cssmin");
		}

		// Load group file
		var groupData = json.parse(fs.readFileSync(groupInName));

		// Reset group directory in git
		shell.exec("cd static/groups/ && touch deleteme && (git ls-files -oz | xargs -0 rm)");

		// Concatenate & md5 groups
		var outData = {js: {}, css: {}};
		var groupOutFiles = {js: {}, css: {}};
		var tmpOut = "";

		for (let type in outData) {
			if (outData.hasOwnProperty(type)) {
				for (let file in groupData[type]) {
					if (groupData[type].hasOwnProperty(file)) {
						var groupName = groupData[type][file];
						var fileName = "static/" + file;

						tmpOut = fs.readFileSync(fileName, "utf8");
						if (minifyGroups) {
							if (type == "css") {
								tmpOut = cssmin(tmpOut);
							} else {
								tmpOut = jsmin(tmpOut);
							}
						}

						if (!outData[type][groupName]) {
							outData[type][groupName] = "";
						}
						outData[type][groupName] += tmpOut + "\n";
					}
				}
			}
		}

		// Rename groups to a unique filename (if they don't match an existing md5)
		for (let type in outData) {
			if (outData.hasOwnProperty(type)) {
				for (let groupName in outData[type]) {
					if (outData[type].hasOwnProperty(groupName)) {
						var fileMd5 = md5(outData[type][groupName]);

						groupOutFiles[type][groupName] = "groups/" + groupName + "-" + fileMd5 + "." + type;
						fs.writeFileSync("static/" + groupOutFiles[type][groupName], outData[type][groupName], {"mode": 0o644});
					}
				}
			}
		}

		// Update the group file with new group file paths
		var groupOut = {js: {}, css: {}};
		for (let type in outData) {
			if (outData.hasOwnProperty(type)) {
				groupOut[type] = {};
				for (let groupName in groupOutFiles[type]) {
					if (groupOutFiles[type].hasOwnProperty(groupName)) {
						groupOut[type][groupName] = groupOutFiles[type][groupName];
					}
				}
			}
		}

		fs.writeFileSync(groupOutName, "var __groups = " + json.stringify(groupOut, null, 2) + ";");
	}

};
