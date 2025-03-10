import { createResource, createSignal } from 'solid-js';
import { LoaderCircle, Play } from 'lucide-solid'
import { Button } from './components/ui/button';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from './components/ui/card';
import { TextFieldTextArea, TextField } from './components/ui/text-field';

export default () => {
  const [inputJson, setInputJson] = createSignal<string>(JSON.stringify([
    {
      "name": "name",
      "type": "string",
      "tag": "{firstname}"
    }
  ]));
  const [data, { refetch }] = createResource(() => fetcher(inputJson()));
  const buttonDisabled = () => inputJson().length < 1 || data?.loading


  const fetchResponse = () => {
    if (data?.loading) {
      return 'Loading...'
    }
    if (data?.error) {
      return data.error
    }
    return data.latest ? JSON.stringify(data.latest, null, 2) : ''
  };

  const handleClear = () => {
    setInputJson('');
  };

  return (
    <div>
      <h1 class='text-center py-4 font-bold text-3xl'>
        Mock Data Generator
      </h1>
      <main class='flex gap-4 justify-center p-2 flex-col md:flex-row'>
        <Card class='grow'>
          <CardHeader>
            <CardTitle>JSON Schema Input</CardTitle>
            <CardDescription>Please enter a properly formatted JSON Schema</CardDescription>
          </CardHeader>
          <CardContent>
            <TextField>
              <TextFieldTextArea
                rows={10}
                value={inputJson()}
                onInput={(e) => setInputJson(e.currentTarget.value)}
              />
            </TextField>
          </CardContent>
          <CardFooter class='flex justify-between'>
            <Button onClick={handleClear}>Clear</Button>
            <Button onClick={refetch} disabled={buttonDisabled()}>
              {data.loading ? (
                <>
                  <LoaderCircle class="animate-spin" /> loading...
                </>
              ) : (
                <>
                  <Play /> Get Mock Data
                </>
              )}
            </Button>
          </CardFooter>
        </Card>

        <Card class='grow'>
          <CardHeader>
            <CardTitle>Generated Mock Data</CardTitle>
            <CardDescription>The mock response data is shown below</CardDescription>
          </CardHeader>
          <CardContent>
            <TextField>
              <TextFieldTextArea
                readOnly
                rows={10}
                value={
                  fetchResponse()
                }
                placeholder='Mock data will appear here'
              />
            </TextField>
          </CardContent>
        </Card>
      </main>
    </div>
  );
};

const fetcher = async (schema: any) => {
  if (!schema) return null;
  const response = await fetch('/api', {
    method: 'POST',
    body: schema,
    signal: AbortSignal.timeout(8000)
  });
  if (!response.ok) throw new Error('Failed to fetch mock data');
  return response.json();
};
