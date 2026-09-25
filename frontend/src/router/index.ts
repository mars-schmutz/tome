import { createRouter, createWebHistory } from "vue-router";
import Overview from "../views/Overview.vue";
import SpellView from "../views/SpellView.vue";

const router = createRouter({
  history: createWebHistory(),
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
