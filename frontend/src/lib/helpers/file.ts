export const readFileAsBytes = (file: File): Promise<number[]> => {
	return new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.onload = () => {
			const arrayBuffer = reader.result as ArrayBuffer;
			const bytes = new Uint8Array(arrayBuffer);
			resolve(Array.from(bytes));
		};
		reader.onerror = reject;
		reader.readAsArrayBuffer(file);
	});
};
