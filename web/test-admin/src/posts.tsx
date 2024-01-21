import { List, Datagrid, TextField, ReferenceField, EditButton } from "react-admin";

export const PostList = () => (
    <List>
        <Datagrid>
            <TextField source="id" />
            <ReferenceField source="userId" reference="users" link="show" />
            <TextField source="title" />
            <EditButton />
        </Datagrid>
    </List>
);