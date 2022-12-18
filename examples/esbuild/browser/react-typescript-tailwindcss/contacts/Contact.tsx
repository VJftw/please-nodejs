import React, { useCallback, useState } from "react";
import { Contact, getContact, isContact } from "./data";

import { Form, useLoaderData } from "react-router-dom";

interface RouteParams {
  contactId: string;
}

export async function loader({ params }: any): Promise<Contact | null> {
  console.log(params);
  return getContact(params.contactId);
}

export default function ContactComponent() {
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

  if (contact == null) {
    console.log("contact is null");
    return <></>;
  }

  return (
    <div className="container mx-auto">
      <div className="">
        <div className="block rounded-xl p-6 sm:p-8">
          <div className="mt-8 mb-0 max-w-md space-y-4">
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
                  className="w-full rounded-lg border-gray-200 p-2 pr-12 text-sm shadow-sm"
                  placeholder="Name"
                  value={contact.name}
                  disabled
                />
              </div>
            </div>

            <div>
              <label htmlFor="notes" className="text-sm font-medium">Notes</label>
              <div className="relative mt-1">
                <textarea
                  id="notes"
                  className="w-full rounded-lg border-gray-200 p-2 pr-12 text-sm shadow-sm min-h-fit h-32"
                  placeholder="Notes"
                  value={contact.notes}
                  disabled
                />
              </div>
            </div>

            <div className="flex items-center justify-between">
              <Form action="edit">
                <button
                  type="submit"
                  className="inline-block rounded-lg bg-blue-500 px-4 py-2 text-sm font-medium text-white"
                >
                  Edit
                </button>
              </Form>
              <Form
                method="post"
                action="delete"
                onSubmit={(event) => {
                  if (
                    !confirm("Please confirm you want to delete this record.")
                  ) {
                    event.preventDefault();
                  }
                }}
                >
                <button
                  type="submit"
                  className="inline-block rounded-lg bg-red-500 px-4 py-2 text-sm font-medium text-white"
                >
                  Delete
                </button>
              </Form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

function Favorite(prop: { contact: Contact }) {
  // yes, this is a `let` for later
  let contact = prop.contact;
  let favorite = contact.favourite;
  return (
    <Form method="post">
      <button
        name="favorite"
        value={favorite ? "false" : "true"}
        aria-label={favorite ? "Remove from favorites" : "Add to favorites"}
      >
        {favorite ? "★" : "☆"}
      </button>
    </Form>
  );
}
