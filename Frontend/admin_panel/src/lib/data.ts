import { User } from "@/types/user";

export const seedUsers: User[] = [
    { id: 1, firstName: 'Ava', lastName: 'Sharma', email: 'ava.sharma@example.com', role: 'Admin', isActive: true },
    { id: 2, firstName: 'Ravi', lastName: 'Kapoor', email: 'ravi.kapoor@example.com', role: 'Editor', isActive: true },
    { id: 3, firstName: 'Neha', lastName: 'Iyer', email: 'neha.iyer@example.com', role: 'Viewer', isActive: false },
    { id: 4, firstName: 'Kabir', lastName: 'Das', email: 'kabir.das@example.com', role: 'Editor', isActive: true },
    { id: 5, firstName: 'Mira', lastName: 'Varma', email: 'mira.varma@example.com', role: 'Viewer', isActive: true },
    { id: 6, firstName: 'Ishaan', lastName: 'Mehta', email: 'ishaan.mehta@example.com', role: 'Admin', isActive: false },
];