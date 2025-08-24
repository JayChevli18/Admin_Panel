import { UserTable } from "@/components/user-table";

const UsersPage=()=>{
    return (
        <div className="space-y-4">
            <h1 className="text-2xl font-semibold tracking-tight">Users</h1>
            <UserTable />
        </div>
    );
}

export default UsersPage;