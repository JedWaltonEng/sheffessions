jest.mock('js-cookie', () => ({
  get: jest.fn(() => null),
  set: jest.fn(),
}));

describe('SubmissionForm Component', () => {
  it('is a placeholder unit test', () => {
    var test = 1;
    expect(test).toBe(1);
  });
});


