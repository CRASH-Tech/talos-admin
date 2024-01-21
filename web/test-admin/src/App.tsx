import { Admin, Resource, ShowGuesser, EditGuesser } from "react-admin";
import { dataProvider } from "./dataProvider";
import { ClusterList } from "./clusters";
import { PostList } from "./posts";
import { UserList } from "./users";

export const App = () => (
  <Admin dataProvider={dataProvider}>
    <Resource name="clusters" list={ClusterList} show={EditGuesser} />
    <Resource name="users" list={UserList} show={ShowGuesser} recordRepresentation="name" />
    <Resource name="posts" list={PostList} edit={EditGuesser} />
  </Admin>

);
