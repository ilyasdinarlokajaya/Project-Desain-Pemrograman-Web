import Vue from 'vue'
import Router from 'vue-router'


// Article
import Article from '../views/Article/list'
import AddArticle from '../views/Article/add'
import EditArticle from '../views/Article/edit'
import DetailArticle from '../views/Article/detail'

// User
import User from '../views/user/list'
import AddUser from '../views/user/add'
import EditUser from '../views/user/edit'
import DetailUser from '../views/user/detail'


// // Notification
// import Notification from '../views/notification/list'
// import AddNotification from '../views/notification/add'
// import EditNotification from '../views/notification/edit'
// import DetailNotification from '../views/notification/detail'

// Product
import Product from '../views/product/list'

Vue.use(Router)

export default new Router({
	routes: [
		// list-article
		{
			path: '/',
			name: 'list-Article',
			component: Article
		},
		{
			path: '/list-Article',
			name: 'list-Article',
			component: Article
		},
		{
			path: '/add-Article',
			name: 'add-Article',
			component: AddArticle
		},
		{
			path: '/edit-Article',
			name: 'edit-Article',
			component: EditArticle
		},
		{
			path: '/detail-Article',
			name: 'detail-Article',
			component: DetailArticle
		},
		// list-user
		{
			path: '/',
			name: 'list-user',
			component: User
		},
		{
			path: '/list-user',
			name: 'list-user',
			component: User
		},
		{
			path: '/add-user',
			name: 'add-user',
			component: AddUser
		},
		{
			path: '/edit-user',
			name: 'edit-user',
			component: EditUser
		},
		{
			path: '/detail-user',
			name: 'detail-user',
			component: DetailUser
		},
		// list-notification
		// {
		// 	path: '/',
		// 	name: 'list-notification',
		// 	component: Notification
		// },
		// {
		// 	path: '/list-notification',
		// 	name: 'list-notification',
		// 	component: Article
		// },
		// {
		// 	path: '/add-notification',
		// 	name: 'add-notification',
		// 	component: AddNotification
		// },
		// {
		// 	path: '/edit-notification',
		// 	name: 'edit-notification',
		// 	component: EditNotification
		// },
		// {
		// 	path: '/detail-notification',
		// 	name: 'detail-notification',
		// 	component: DetailNotification
		// },
		{
			path: '/list-product',
			name: 'list-product',
			component: Product
		}
	]
})
