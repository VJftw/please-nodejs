import React, { useCallback, useState } from "react";
import { Form, useLoaderData , redirect, useNavigate} from "react-router-dom";
import { Contact, updateContact, isContact } from "./data";



export default function EditContact() {
  const navigate = useNavigate();
  const [contact, setContact] = useState<Contact>({
    avatar: "",
    createdAt: Date.now(),
    favourite: false,
    id: "",
    name: "",
    notes: "",
  });

  const loadedContact = useLoaderData();
  if (isContact(loadedContact)) {
    if (loadedContact != contact) {
      setContact(loadedContact);
    }
  }

  return (
    <div className="container mx-auto">
      <div className="">
        <div className="block rounded-xl p-6 sm:p-8">
          <Form method="post" id="contact-form" className="mt-8 mb-0 max-w-md space-y-4">
            <div>
              <label htmlFor="name" className="text-sm font-medium">Name</label>
              <div className="relative mt-1 flex flex-row">
                <div className="p-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 text-yellow-400"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                </div>
                <input
                  type="text"
                  id="name"
                  name="name"
                  className="w-full rounded-lg border-gray-200 p-2 pr-12 text-sm shadow-sm"
                  placeholder="Name"
                  defaultValue={contact.name}
                />
              </div>
            </div>

            <div>
              <label htmlFor="notes" className="text-sm font-medium">Notes</label>
              <div className="relative mt-1">
                <textarea
                  id="notes"
                  name="notes"
                  className="w-full rounded-lg border-gray-200 p-2 pr-12 text-sm shadow-sm min-h-fit h-32"
                  placeholder="Notes"
                  defaultValue={contact.notes}
                />
              </div>
            </div>

            <div className="flex items-center justify-between">
                <button type="submit"
                  className="inline-block rounded-lg bg-blue-500 px-4 py-2 text-sm font-medium text-white"
                >
                  Save
                </button>
                <button
                  type="button"
                  onClick={() => {
                    navigate(-1);
                  }}
                  className="inline-block rounded-lg bg-green-500 px-4 py-2 text-sm font-medium text-white"
                >
                  Cancel
                </button>
            </div>
          </Form>
        </div>
      </div>
    </div>
  );
}
