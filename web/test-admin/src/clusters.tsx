import { Create, SimpleForm, TextInput, List, Datagrid, TextField, ReferenceField, EditButton, required } from "react-admin";

export const ClusterList = () => (
    <List>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <EditButton />
        </Datagrid>
    </List>
);

export const ClusterCreate = () => (
    <Create>
        <SimpleForm>
            <TextInput source="name" validate={[required()]} fullWidth />
        </SimpleForm>
    </Create>
);