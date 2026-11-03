import { createRouter, createWebHashHistory } from "vue-router";
import Overview from "../views/Overview.vue";
import SpellView from "../views/SpellView.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      component: Overview,
    },
    {
      path: "/spells",
      component: SpellView,
    },
  ],
});

export default router;
