package cmd

func InitializeProject(config *Config) {
	// --- Initialization ---
	// 1. Copy nextjs-base to desired location & rename it to config.ProjectName.
	// --- Pages & their components ---
	// 2. Create files for each page in project/pages
	//    2.1. Get "common/page.jsx", rename it and use it.
	//    2.2. For each Page.Component:
	//         2.2.1. Find the component's folder in "components/componentName"
	//         2.2.2. We will have files of different type in these folders. Each one ends differently: _api, _util etc.
	//                This suffix points to where this components belongs. Copy them to their folders inside of the project.
	//         2.2.3: These components share same dir-structure, so they all are already connected theoretically,
	//                But we still have to connect them to their pages! (only _component's)
	//				  Each page has comments specifically for this occasion:
	//                // --- imports start --- ,  --- imports end ---
	//                // --- content start --- , --- content end ---
	//				  Find these comments, import & declare Page's components.
	// --- Theming ---
	// 3. Grab required theme from /themes and put it in project/styles/theme.js (emotionCSS theming, optional dark mode)
	//
	// *Notes:
	// Parse jsx files using goquery.
}
