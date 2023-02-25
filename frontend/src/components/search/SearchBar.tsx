import { TextInput } from "@mantine/core";
import { IconSearch } from "@tabler/icons";
import { type Dispatch, type SetStateAction } from "react";

type SearchProps = {
	value: string;
	setValue: Dispatch<SetStateAction<string>>;
};
export function SearchBar({ value, setValue }: SearchProps) {
	return (
		<TextInput
			value={value}
			onChange={(e) => setValue(e.currentTarget.value)}
			placeholder="Search"
			icon={<IconSearch size={16} />}
			radius={0}
			size="md"
			w="100%"
			autoComplete="off"
			style={{ zIndex: 1 }}
		/>
	);
}
