import { redirect } from "react-router-dom";
import { deleteContact, isContact, updateContact} from "./data";

export async function deleteAction({ params }: any) {
  await deleteContact(params.contactId);
  return redirect("/");
}

export async function editAction({ request, params }: any) {
    const formData = await request.formData();
    const updates = Object.fromEntries(formData);
    if (isContact(updates)) {
      await updateContact(params.contactId, updates);
      return redirect(`/contacts/${params.contactId}`);
    }
  }
