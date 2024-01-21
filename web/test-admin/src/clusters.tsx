import { List, Datagrid, TextField, ReferenceField, EditButton } from "react-admin";

export const ClusterList = () => (
    <List>
        <Datagrid>
            <TextField source="id" />
            <TextField source="title" />
            <EditButton />
        </Datagrid>
    </List>
);