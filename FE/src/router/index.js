import Vue from 'vue'
import Router from 'vue-router'

// User
import User from '../views/user/list'
import AddUser from '../views/user/add'
import EditUser from '../views/user/edit'
import DetailUser from '../views/user/detail'

// Product
import Product from "../views/product/list";
import AddProduct from "../views/product/add";
import DetailProduct from "../views/product/detail";
import EditProduct from "../views/product/edit";

// Voucher
import Voucher from "../views/voucher/list";
import Addvoucher from "../views/voucher/add";
import Editvoucher from "../views/voucher/edit";
import Detailvoucher from "../views/voucher/detail";

// Banner
import Banner from '../views/banner/list';
import AddBanner from "../views/banner/add";
import EditBanner from "../views/banner/edit";
import DetailBanner from "../views/banner/detail";

// Plannogram
import Plannogram from '../views/plannogram/list';
import AddPlannogram from "../views/plannogram/add";
import EditPlannogram from "../views/plannogram/edit";
import DetailPlannogram from "../views/plannogram/detail";

// Article
import Article from "../views/article/list";
import AddArticle from "../views/article/add";
import EditArticle from "../views/article/edit";
import DetailArticle from "../views/article/detail";

// Notification
import Notification from '../views/notification/list';
import AddNotification from "../views/notification/add";
import EditNotification from "../views/notification/edit";
import DetailNotification from "../views/notification/detail";

// DeliveryMethod
import DeliveryMethod from '../views/delivery-method/list';
import AddDeliveryMethod from "../views/delivery-method/add";
import EditDeliveryMethod from "../views/delivery-method/edit";
import DetailDeliveryMethod from "../views/delivery-method/detail";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: "/",
      name: "list-user",
      component: User
    },
    {
      path: "/list-user",
      name: "list-user",
      component: User
    },
    {
      path: "/add-user",
      name: "add-user",
      component: AddUser
    },
    {
      path: "/edit-user/:id",
      name: "edit-user",
      component: EditUser
    },
    {
      path: "/detail-user/:id",
      name: "detail-user",
      component: DetailUser
    },
    {
      path: "/list-product",
      name: "list-product",
      component: Product
    },
    {
      path: "/add-product",
      name: "add-product",
      component: AddProduct
    },
    {
      path: "/detail-product/:id",
      name: "detail-product",
      component: DetailProduct
    },
    {
      path: "/edit-product/:id",
      name: "edit-product",
      component: EditProduct
    },
    {
      path: "/list-voucher",
      name: "list-voucher",
      component: Voucher
    },
    {
      path: "/add-voucher",
      name: "add-voucher",
      component: Addvoucher
    },
    {
      path: "/edit-voucher/:id",
      name: "edit-voucher",
      component: Editvoucher
    },
    {
      path: "/detail-voucher/:id",
      name: "detail-voucher",
      component: Detailvoucher
    },
    {
      path: "/list-banner",
      name: "list-banner",
      component: Banner
    },
    {
      path: "/add-banner",
      name: "add-banner",
      component: AddBanner
    },
    {
      path: "/edit-banner/:id",
      name: "edit-banner",
      component: EditBanner
    },
    {
      path: "/detail-banner/:id",
      name: "detail-banner",
      component: DetailBanner
    },
    {
      path: "/list-plannogram",
      name: "list-plannogram",
      component: Plannogram
    },
    {
      path: "/add-plannogram",
      name: "add-plannogram",
      component: AddPlannogram
    },
    {
      path: "/edit-plannogram/:id",
      name: "edit-plannogram",
      component: EditPlannogram
    },
    {
      path: "/detail-plannogram/:id",
      name: "detail-plannogram",
      component: DetailPlannogram
    },
    {
      path: "/list-article",
      name: "list-article",
      component: Article
    },
    {
      path: "/add-article",
      name: "add-article",
      component: AddArticle
    },
    {
      path: "/edit-article/:id",
      name: "edit-article",
      component: EditArticle
    },
    {
      path: "/detail-article/:id",
      name: "detail-article",
      component: DetailArticle
    },
    {
      path: "/list-notification",
      name: "list-notification",
      component: Notification
    },
    {
      path: "/add-notification",
      name: "add-notification",
      component: AddNotification
    },
    {
      path: "/edit-notification/:id",
      name: "edit-notification",
      component: EditNotification
    },
    {
      path: "/detail-notification/:id",
      name: "detail-notification",
      component: DetailNotification
    },
    {
      path: "/list-delivery-method",
      name: "list-delivery-method",
      component: DeliveryMethod
    },
    {
      path: "/add-delivery-method",
      name: "add-delivery-method",
      component: AddDeliveryMethod
    },
    {
      path: "/edit-delivery-method/:id",
      name: "edit-delivery-method",
      component: EditDeliveryMethod
    },
    {
      path: "/detail-delivery-method/:id",
      name: "detail-delivery-method",
      component: DetailDeliveryMethod
    }
  ]
});
