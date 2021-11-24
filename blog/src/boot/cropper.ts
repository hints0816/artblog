import VueCropper from 'vue-cropper'; 
import 'vue-cropper/dist/index.css'
import { boot } from 'quasar/wrappers';

export default boot(({ app }) => {
    app.use(VueCropper)
});
  