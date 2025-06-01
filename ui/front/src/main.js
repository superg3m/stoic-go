import './assets/main.css';

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';

import Aura from '@primeuix/themes/aura';

import Menubar from 'primevue/menubar'
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Dialog from 'primevue/dialog';
import Card from 'primevue/card';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Checkbox from 'primevue/checkbox';
import RadioButton from 'primevue/radiobutton';
import ToggleButton from 'primevue/togglebutton';
import Textarea from 'primevue/textarea';
import Toast from 'primevue/toast';
import ProgressSpinner from 'primevue/progressspinner';
import Tooltip from 'primevue/tooltip';
import Avatar from 'primevue/avatar';
import Badge from 'primevue/badge';
import Drawer from 'primevue/drawer';

import 'primeicons/primeicons.css'

const app = createApp(App);

app.use(router);
app.use(PrimeVue, {
  theme: {
    preset: Aura
  }
});
app.use(ToastService);

app.component('Menubar', Menubar);
app.component('Button', Button);
app.component('InputText', InputText);
app.component('Dialog', Dialog);
app.component('Card', Card);
app.component('DataTable', DataTable);
app.component('Column', Column);
app.component('Checkbox', Checkbox);
app.component('RadioButton', RadioButton);
app.component('ToggleButton', ToggleButton);
app.component('Textarea', Textarea);
app.component('Toast', Toast);
app.component('ProgressSpinner', ProgressSpinner);
app.component('Avatar', Avatar);
app.component("Badge", Badge);
app.component("Drawer", Drawer);

app.directive('tooltip', Tooltip);

app.mount('#app');
