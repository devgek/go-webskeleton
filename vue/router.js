// Routes
const routes = [
  {
    path: "/login",
    component: gekLoginView,
    props: {mainHeader: 'Bitte anmelden'}
  },
  {
    path: '/',
    component: gekLayoutView,
    children: [
      {
        path: 'start.html',
        alias: '',
        component: gekHomeView,
        name: 'Start',
        meta: {description: 'start app'}
      }, {
        path: 'home',
        component: gekHomeView,
        name: 'Home',
        meta: {description: 'show home'}
      }, {
        path: 'page1',
        component: gekPage1View,
        name: 'Page1',
        meta: {description: 'show page1'}
      }
    ]
  },
  {
    // not found handler
    path: "*",
    component: gekNotFoundView,
  },
];

const router = new VueRouter({
  routes,
});

router.beforeEach((to, from, next) => {
    // redirect to login page if user is not logged in and trying to access a restricted page
    const publicPages = ['/', '/login', '/page2']
    const authRequired = !publicPages.includes(to.path)
    const loggedIn = localStorage.getItem('user')
  
    if (authRequired && !loggedIn) {
      return next('/login')
    }
  
    next()
  })