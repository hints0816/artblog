import { boot } from 'quasar/wrappers';
import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';

import VMdPreviewHtml from '@kangc/v-md-editor/lib/preview-html';
import '@kangc/v-md-editor/lib/style/preview-html.css';

import Prism from 'prismjs';

import VMdPreview from '@kangc/v-md-editor/lib/preview';
import '@kangc/v-md-editor/lib/style/preview.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';

// highlightjs
import hljs from 'highlight.js';


export default boot(({ app }) => {
  VMdPreview.use(githubTheme, {
    Hljs: hljs,
  });
  VueMarkdownEditor.use(vuepressTheme, {
    Prism,
  })
  app.use(VueMarkdownEditor);
  app.use(VMdPreviewHtml);
  app.use(VMdPreview);
});
