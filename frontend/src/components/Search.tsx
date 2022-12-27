import { TextInput } from "@mantine/core";
import { IconSearch } from "@tabler/icons";
import { Dispatch, SetStateAction } from "react";

interface SearchProps {
	value: string;
	setValue: Dispatch<SetStateAction<string>>;
}
export const Search = ({ value, setValue }: SearchProps) => {
	return (
		<TextInput
			value={value}
			onChange={(e) => setValue(e.currentTarget.value)}
			placeholder="Search"
			icon={<IconSearch size={16} />}
			radius={0}
			size="md"
			w="100%"
			pos="fixed"
			autoComplete="off"
			style={{ zIndex: 1 }}
		/>
	);
};
